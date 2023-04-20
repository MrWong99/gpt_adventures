# Creating new Objects with Gamestate

Since there are no boundaries about the object a character wants to search for or interact with, we need to be able to create new objects on-the-fly.
This is done by creating a new object with a finite-state machine that allowes the object to be rendered with a limited set of images while also limiting a users interactions to a managable amount.

## Initialization Prompt

```
You are an agent tasked to generate a Finite-state machine for an object. This object will be integrated into a turn based adventure game and therefore the states and the transitions between states that you generate should be typical for a game. You will also be given a users intended use for the object which you should take into consideration but if the intended use makes no sense for the given object you can ignore it.

You are given the following JSON-Object format as input:
{
"INTENT": "the users intent on how to use the object",
"OBJECT": "the name of the object",
"DESCRIPTION": "a short description of the object"
}


You must enrich this input with a Finite-state machine and output it in this JSON-Output format:
{
"INTENT": "the users intent on how to use the object",
"OBJECT": "the name of the object",
"DESCRIPTION": "a short description of the object",
"STATE": "the initial state of the object",
"AVAILABLE_STATES": {
"S_1": "the first available state",
"S_1_TRANSITIONS": {
"T_S_1": "the target state of the first transition starting at the first available state",
"T_1_NAME": "the name of the first transition operation for the first state",
"T_1_REQUIREMENT": "one or many optionally required objects that would be needed to perform the second transition",
"T_S_2": "the target state of the second transition starting at the first available state",
"T_2_NAME": "the name of the second transition operation for the first state",
"T_2_REQUIREMENT": "one or many optionally required objects that would be needed to perform the second transition",
...
},
"S_2": "the second available state",
"S_2_TRANSITIONS": {
"T_S_1": "the target state of the first transition starting at the second available state",
"T_1_NAME": "the name of the first transition operation for the second state",
"T_1_REQUIREMENT": "one or many optionally required objects that would be needed to perform the second transition",
"T_S_2": "the target state of the second transition starting at the first available state",
"T_2_NAME": "the name of the second transition operation for the second state",
"T_2_REQUIREMENT": "one or many optionally required objects that would be needed to perform the second transition",
...
},
...
}

For example:
{
"INTENT": "lock so I can sleep in peace at night",
"OBJECT": "door",
"DESCRIPTION": "the door leading from the basement to the hallway",
"STATE": "CLOSED_UNLOCKED",
"AVAILABLE_STATES": {
"S_1": "OPENED",
"S_1_TRANSITIONS": {
"T_S_1": "CLOSED_UNLOCKED",
"T_1_NAME": "close",
"T_1_REQUIREMENT": ""
},
"S_2": "CLOSED_UNLOCKED",
"S_2_TRANSITIONS": {
"T_S_1": "OPENED",
"T_1_NAME": "open",
"T_S_2": "CLOSED_LOCKED",
"T_2_NAME": "lock",
"T_2_REQUIREMENT": "key"
},
"S_3": "CLOSED_LOCKED",
"S_3_TRANSITIONS": {
"T_S_1": "CLOSED_UNLOCKED",
"T_1_NAME": "unlock",
"T_1_REQUIREMENT": "key"
},
"S_3": "BROKEN",
"S_3_TRANSITIONS": {
"T_S_1": "OPEN",
"T_1_NAME": "repair",
"T_1_REQUIREMENT": "repair kit and WD40 spray"
}
}
}
```

## Example Input

```
This is the next JSON-Object input:
{
"INTENT": "clean myself",
"OBJECT": "shower",
"DESCRIPTION": "the shower in the first floor bathroom"
}


Provide the fitting JSON-Output response that now has a Finite-state machine.
```

## Example Output

```
{
"INTENT": "clean myself",
"OBJECT": "shower",
"DESCRIPTION": "the shower in the first floor bathroom",
"STATE": "OFF",
"AVAILABLE_STATES": {
"S_1": "OFF",
"S_1_TRANSITIONS": {
"T_S_1": "ON",
"T_1_NAME": "turn on",
"T_1_REQUIREMENT": "water supply and functional faucet"
},
"S_2": "ON",
"S_2_TRANSITIONS": {
"T_S_1": "OFF",
"T_1_NAME": "turn off",
"T_1_REQUIREMENT": "",
"T_S_2": "BROKEN",
"T_2_NAME": "break",
"T_2_REQUIREMENT": "heavy object"
},
"S_3": "BROKEN",
"S_3_TRANSITIONS": {
"T_S_1": "OFF",
"T_1_NAME": "repair",
"T_1_REQUIREMENT": "plumber and replacement parts"
}
}
}
```

## Error Handling

This doesn't need any error handling since an invalid intent will just be ignored during creation of the Finite-state machine.
The invalid intent will however result in error handling by the [mapping tool for user input](/agent_concepts/object_interactions_mapping.md).
