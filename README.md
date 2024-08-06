# Image-to-ASCII

## Description

The image-to-ASCII program converts PNG and JPG images into ASCII art. This tool allows you to transform your images into a text-based representation using various ASCII characters. You can also adjust the pixel step frequency to control the level of detail in the ASCII output.


## Usage
To convert an image to ASCII, use the following command:
```bash
./program image.png image.txt
```
### Where:
- image.png is the path to your input image file (PNG or JPG).
- image.txt is the path where the ASCII art output will be saved.

### Optional: Step Frequency
You can specify a step frequency to control the level of detail. The default is 2. Adding a number at the end of the command adjusts the frequency of the pixels processed. For example:
```bash
./program image.png image.txt 3
```

## Example
Here is an example of a standard output 1/4 size 1192x1590 image:

<img src="example.png" width=500>
