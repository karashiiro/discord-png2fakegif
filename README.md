# discord-png2fakegif
Converts Discord emote PNGs into (mostly) still GIFs that Discord recognizes as animated emotes.

Just converting a PNG into a GIF does not at this time allow the emote to be stored in an animated emote slot, because they aren't actually animated.
This program creates a GIF with two frames, one with a single pixel changed from the other, that does count as an animated emote.

This process saturates the emote a bit. Trying the two color palettes Go offers resulted in one looking washed-out and the other looking saturated,
so I chose to keep the saturated one.

This will lock the emote to Nitro users only. Most servers don't need this.

## Usage
`discord-png2fakegif.exe ./emote.png`

Or, just drop the PNG on the executable. The output file will be named `emote.png.gif` if your input filename is `emote.png`.

## Example

### Input PNG
![Example input](./example.png)

### Output GIF
![Example output](./example.gif)

### Uploading
Uploading the output to Discord shows that it interprets the GIF as animated.
![Example upload](./upload_example.png)
