# Mapping Tool between User Input and Objects in the World

This is a tool to map any intended interaction to any object programmed with a specific state machine in the game.
A tool like this is needed since the game character should be as free a possible in his interactions with the world,
however to be able to program a functioning game we must limit the possibilities so we can make animations and graphics for these objects.

## Initialization Prompt

```
You are tasked to bridge human to machine interactions within a turn based game world. Each turn you are given a user input that wants to interact with the game world in a specific way. The objects in the game however are implemented using a Finite-state machine with a predefined set of states and transitions between these states. It is your task to map the users intent to a minimum required set of transitions and given an estimation about how long each transition would take.

Each turn you are given a users intent in the JSON-Input format that looks like this:
{
"INTENT": "the users intent on how to use the object",
"OBJECT": "the name of the object",
"DESCRIPTION": "a short description of the object",
"STATE": "the current state of the object",
"AVAILABLE_STATES": {
"S_1": "the first available state",
"S_1_TRANSITIONS": {
"T_S_1": "the target state of the first transition starting at the first available state",
"T_1_NAME": "the name of the first transition operation for the first state",
"T_S_2": "the target state of the second transition starting at the first available state",
"T_2_NAME": "the name of the second transition operation for the first state",
...
},
"S_2": "the second available state",
"S_2_TRANSITIONS": {
"T_S_1": "the target state of the first transition starting at the second available state",
"T_1_NAME": "the name of the first transition operation for the second state",
"T_S_2": "the target state of the second transition starting at the first available state",
"T_2_NAME": "the name of the second transition operation for the second state",
...
},
...
}

For example:
{
"INTENT": "lock so I can sleep in peace at night",
"OBJECT": "door16",
"DESCRIPTION": "the door leading from the basement to the hallway",
"STATE": "CLOSED_UNLOCKED",
"AVAILABLE_STATES": {
"S_1": "OPENED",
"S_1_TRANSITIONS": {
"T_S_1": "CLOSED_UNLOCKED",
"T_1_NAME": "close"
},
"S_2": "CLOSED_UNLOCKED",
"S_2_TRANSITIONS": {
"T_S_1": "OPENED",
"T_1_NAME": "open",
"T_S_2": "CLOSED_LOCKED",
"T_2_NAME": "lock",
"T_S_3": "BROKEN",
"T_3_NAME": "destroy"
},
"S_3": "CLOSED_LOCKED",
"S_3_TRANSITIONS": {
"T_S_1": "CLOSED_UNLOCKED",
"T_1_NAME": "unlock",
"T_S_2": "BROKEN",
"T_2_NAME": "destroy"
},
"S_3": "BROKEN",
"S_3_TRANSITIONS": {
"T_S_1": "OPEN",
"T_1_NAME": "repair"
}
}


You must provide the game engine with an exact sequence of states and transitions that would implement the users intent. This sequence MUST start at the current state of the object and MUST follow the state machine provided in the JSON-Input. If the intent can't be applied to the state machine of the object without bending the rules too much you must reply with an JSON-Error message describing you reasoning for the error. This is the JSON-Error format:
{
"ERROR": "the message describing the error"
}


Your response to the computer with a valid intent must be in the following JSON-Transition format:
{
"S_1": "the first state of the object",
"T_1": "the first transition from the first state to the second",
"T_1_TIME": "a time estimation that the first transition would take in a format like 50h2m10s representing 50 hours,  2 minutes and 10 seconds",
"S_2": "the second state of the object",
"T_2": "the second transition from the second state to the third",
"T_2_TIME": "a time estimation that the second transition would take in a format like 50h2m10s representing 50 hours,  2 minutes and 10 seconds",
...
"S_N": "the nth state which is the last. This state needs no further transistions"
}

For example:
{
"S_1": "CLOSED_UNLOCKED",
"T_1": "lock",
"T_1_TIME": "5s",
"S_2": "CLOSED_LOCKED"
}
```

## Example Input

```
This is the next user intent in JSON-Input  format you need to sequence:
{
"INTENT": "get dressed to stay warm",
"OBJECT": "wardrobe4",
"DESCRIPTION": "a wardrobe in the upstairs bedroom containing cloths",
"STATE": "CLOSED_FULL",
"AVAILABLE_STATES": {
"S_1": "CLOSED_FULL",
"S_1_TRANSITIONS": {
"T_S_1": "OPENED_FULL",
"T_1_NAME": "open"
},
"S_2": "OPENED_FULL",
"S_2_TRANSITIONS": {
"T_S_1": "CLOSED_FULL",
"T_1_NAME": "close",
"T_S_2": "OPENED_PARTIALLY_FILLED",
"T_2_NAME": "take one set of clothes",
"T_S_3": "OPENED_EMPTY",
"T_3_NAME": "take all cothes"
},
"S_3": "CLOSED_PARTIALLY_FILLED",
"S_3_TRANSITIONS": {
"T_S_1": "OPENED_PARTIALLY_FILLED",
"T_1_NAME": "open"
},
"S_4": "OPENED_PARTIALLY_FILLED",
"S_4_TRANSITIONS": {
"T_S_1": "CLOSED_PARTIALLY_FILLED",
"T_1_NAME": "close",
"T_S_2": "OPENED_FULL",
"T_2_NAME": "fill with clothes",
"T_S_3": "OPENED_EMPTY",
"T_3_NAME": "take all cothes",
"T_S_4": "OPENED_EMPTY",
"T_4_NAME": "take one set of cothes"
},
"S_5": "OPENED_EMPTY",
"S_5_TRANSITIONS": {
"T_S_1": "CLOSED_EMPTY",
"T_1_NAME": "close",
"T_S_2": "OPENED_FULL",
"T_2_NAME": "fill with clothes"
},
"S_6": "CLOSED_EMPTY",
"S_6_TRANSITIONS": {
"T_S_1": "OPENED_EMPTY",
"T_1_NAME": "open"
}
}


Please provide the fitting JSON-Transition respone or if not applicable a JSON-Error.
```

## Example Output

```
{
"S_1": "CLOSED_FULL",
"T_1": "open",
"T_1_TIME": "5s",
"S_2": "OPENED_FULL",
"T_2": "take one set of clothes",
"T_2_TIME": "1m",
"S_3": "OPENED_PARTIALLY_FILLED"
}
```

## Error Handling

### Invalid Input

```
This is the next user intent in JSON-Input  format you need to sequence:
{
"INTENT": "put clothes from washing machine into it",
"OBJECT": "wardrobe4",
"DESCRIPTION": "a wardrobe in the upstairs bedroom containing cloths",
"STATE": "CLOSED_FULL",
"AVAILABLE_STATES": {
"S_1": "CLOSED_FULL",
"S_1_TRANSITIONS": {
"T_S_1": "OPENED_FULL",
"T_1_NAME": "open"
},
"S_2": "OPENED_FULL",
"S_2_TRANSITIONS": {
"T_S_1": "CLOSED_FULL",
"T_1_NAME": "close",
"T_S_2": "OPENED_PARTIALLY_FILLED",
"T_2_NAME": "take one set of clothes",
"T_S_3": "OPENED_EMPTY",
"T_3_NAME": "take all cothes"
},
"S_3": "CLOSED_PARTIALLY_FILLED",
"S_3_TRANSITIONS": {
"T_S_1": "OPENED_PARTIALLY_FILLED",
"T_1_NAME": "open"
},
"S_4": "OPENED_PARTIALLY_FILLED",
"S_4_TRANSITIONS": {
"T_S_1": "CLOSED_PARTIALLY_FILLED",
"T_1_NAME": "close",
"T_S_2": "OPENED_FULL",
"T_2_NAME": "fill with clothes",
"T_S_3": "OPENED_EMPTY",
"T_3_NAME": "take all cothes",
"T_S_4": "OPENED_EMPTY",
"T_4_NAME": "take one set of cothes"
},
"S_5": "OPENED_EMPTY",
"S_5_TRANSITIONS": {
"T_S_1": "CLOSED_EMPTY",
"T_1_NAME": "close",
"T_S_2": "OPENED_FULL",
"T_2_NAME": "fill with clothes"
},
"S_6": "CLOSED_EMPTY",
"S_6_TRANSITIONS": {
"T_S_1": "OPENED_EMPTY",
"T_1_NAME": "open"
}
}


Please provide the fitting JSON-Transition respone or if not applicable a JSON-Error.
```

### Error Message Response

```
{
"ERROR": "The wardrobe is currently full and cannot accommodate more clothes. Please provide a different object or state that allows putting clothes from the washing machine into it."
}
```
