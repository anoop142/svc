# svc
## Simple Video Compressor 


A super simple video compressor using FFMPEG 
optimised for compressing online class / screen recordings.

It is more like a wrapper .

Originally written for teachers.

FFMPEG for encoding .
Mediainfo for checking resolution.

## Features
* Simple, very little input from user.
* Encoding using x264 .
* Web compatible output.
* CRF based encoding, hence faster than 2-pass and better qulaity than single pass.
* Optimised encoding  depending on the source.
* Maximum reduction in file size without breaking 
video playback in android devices.
* Default optimised for screen reordings.
* Auto 720p downscale option.
* Batch mode.
* Audio compression using opus.

## Requirements
ffmpeg

mediainfo

## Usage

1. place video files to compress in Input folder
2. run compressor & choose crf value
3. done

For Windows  put ffmpeg  and mediainfo binaries in bin folder.