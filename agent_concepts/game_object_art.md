# Create new Game Objects on-the-fly

If a user requests to interact or search for an object that is not yet defined for the game world, the agent should be able to create a new object on-the-fly.
To do this we create a simple prompt for DALL-E to generate the first image and then adapt this image to display the different states of the object.

## Prompt

```
<first object state> <object>, isometric, pixel art
```

## Example

```
closed door, isometric, pixel art
```
