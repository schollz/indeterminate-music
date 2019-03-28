import operator
import os.path
import random
import json

# pip install pretty_midi
import pretty_midi

chords = json.load(open('data/chords.json'))
chord_types = json.load(open('data/chord_types.json'))
c_major_scale = ("C D E F G A B " * 8).split()
c_chromatic_scale = ("C Db D Eb E F F# G Ab A Bb B " * 20).split()


def get_midi_notes(notes, register=4):
    """Get notes like
    ['C','G']
    And return the midi notes corresponding to
    C4 G4, which is
    [48, 55]
    Optional: Change register
    """
    midi_notes = []
    for i, note in enumerate(notes):
    	midi_note = register*12 + c_chromatic_scale.index(note)
    	if i > 0:
    		while midi_note < midi_notes[i-1]:
	    		midi_note += 12
    	midi_notes.append(midi_note)
    return midi_notes


def get_chord(chord_string, root='C'):
    """Get a chord string like
    1 3 5 7b
    And return the notes in the chord
    C E G Bb
    Optional: Change root
    """
    notes = []
    for i, c in enumerate(chord_string.split()):
        num = int(c.replace('b', '').replace('#', '')) - 1
        num_chro = c_chromatic_scale.index(c_major_scale[num])
        if 'b' in c:
            num_chro -= 1
        if '#' in c:
            num_chro += 1
        if root != 'C':
            num_chro += c_chromatic_scale.index(root)
        notes.append(c_chromatic_scale[num_chro])
    return notes


def notes_to_chord(note_array, enforce_root=False):
    """Get a note string like
    ['C','E','G','Bb']
    And return the chord
    C7
    Optional: enforce root
    """
    notes = note_array
    matches = {}
    for c in all_chords:
        matches[c] = len(set(all_chords[c]) & set(notes))
        if len(set(all_chords[c])) > len(set(notes)):
            matches[c] = matches[c] / len(set(all_chords[c]))
        else:
            matches[c] = matches[c] / len(set(notes))

    best_match = "0" * 100
    best_score = 0
    for i, match in enumerate(sorted(matches.items(), key=operator.itemgetter(1), reverse=True)):
        if i == 0:
            best_score = match[1]
        if match[1] < best_score and best_match != "0" * 100:
            break
        if len(match[0]) < len(best_match):
            if enforce_root:
                if notes[0] == all_chords[match[0]][0]:
                    best_match = match[0]
            else:
                best_match = match[0]
    return best_match


def chord_to_notes(chord_string, voicing=False, preserve_root=True):
	"""Get a chord like
	C7
	and return the notes
	['C','E','G','Bb']

	optionally you can allow different voicings
	"""
	chord = all_chords[chord_string]
	if voicing:
		if preserve_root:
			root = chord[0]
			other_notes = chord[1:]
			random.shuffle(other_notes)
			chord = [root] + other_notes
		else:
			random.shuffle(chord)
	return chord


def generate_all_chords():
    all_chords = {}
    for c in chord_types['simple']:
        for note in c_chromatic_scale:
            all_chords[note + c] = get_chord(chords[c], root=note)

    with open('data/all_chords.json', 'w') as f:
        f.write(json.dumps(all_chords, indent=2))

def midi_to_note(midi_value):
    return c_chromatic_scale[midi_value]

if not os.path.isfile('data/all_chords.json'):
    generate_all_chords()
all_chords = json.load(open('data/all_chords.json', 'r'))


midi_data = pretty_midi.PrettyMIDI('testing/toto-africa.mid')
time_per_measure = 0
downbeats = midi_data.get_downbeats()
print(downbeats)

downbeatsHalf = []
for i,downbeat in enumerate(downbeats):
    if i==0:
        downbeatsHalf.append(downbeat)
        continue 
    downbeatsHalf.append((downbeat-downbeats[i-1])/2+downbeats[i-1])
    downbeatsHalf.append(downbeat)
downbeats = downbeatsHalf
    
    

for instrument_i,instrument in enumerate(midi_data.instruments):
    if instrument.is_drum:
        continue
    measures = [[]]*len(downbeats)
    
    # add a note for each measure that it is played
    for note in instrument.notes:
        noteStart = note.start +0.0001
        noteEnd = note.end - 0.001
        for i,_ in enumerate(downbeats):
            if i ==0:
                continue
            if noteStart < downbeats[i]:
                if len(measures[i-1]) == 0:
                    measures[i-1] = []
                measures[i-1].append(note)
            if noteEnd < downbeats[i]:
                break

    chords = []
    for i,measure in enumerate(measures):
        notes = []
        for note in measure:
            notes.append(midi_to_note(note.pitch))
        if len(notes) > 1:
            chords.append(notes_to_chord(notes).ljust(5))
        else:
            chords.append("     ")
    print(str(instrument_i) + ": " + " ".join(chords))

times = []
for _, downbeat in enumerate(downbeats):
    times.append("{:2.1f}".format(downbeat).ljust(5))
print("t: " + " ".join(times))