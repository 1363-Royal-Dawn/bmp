### DIMENSIONS

* LinkedIn
    * Logo 320x132
    * Cover Image 1128x191
* Favicon
    * 16x16
* Youtube
    * Banner #1 2048x1152
    * Banner #2 2560x1440
    * Watermark 150x150
* Discord
  * 512x512
* Google
  * Workspace logo 320x132 PNG/GIF
    

# FILE FORMATS
The icon file-format follows a simple and strict layout, remember that it is very tightly connected with Bitmaps.

## ICON HEADER
6 Bytes

* 0x00 RESERVED
* 0x02 TYPE (0 EQ CURSOR, 1 EQ ICON)
* 0x04 COUNT 

## ICON DIRECTORY
Directory 16 bytes

0x00 1 Width, 0 == 256, n < 256

0x01 1 Height, 0 == 256, n < 256

0x02 1 ColorCount, 0 >= 256, n < 256

0x03 1 Reserved

0x04 2 Color_XHotSpot, 0 OR 1 OR THE X IF CURSOR

0x06 2 BitsPerPixel_YHotSpot, OR THE Y IF CURSOR

0x08 4 Data in Bytes

0x0C 4 Offset in file


## BITMAP INFO HEADER
For the bitmap info header, we only rely on the first few bytes.
BITMAPINFOHEADER 40 BYTES (ONLY )

0x00 SIZE

0x04 WIDTH

0x08 HEIGHT (2*HEIGHT)

0x0C COLOR_PLANE (ALWAYS 1)

0x0E BITCOUNT 

0x10 FILL 24 BYTES WITH ZERO

0x28 IMAGE DATA (DEPENDS ON THE BITCOUNT, FOR NOW 32 BIT (RGBA) AS PIXELDATA)



### BIT COUNT

```
BI_BITCOUNT_0 = 0x0000,
BI_BITCOUNT_1 = 0x0001,
BI_BITCOUNT_2 = 0x0004,
BI_BITCOUNT_3 = 0x0008,
BI_BITCOUNT_4 = 0x0010,
BI_BITCOUNT_5 = 0x0018,
BI_BITCOUNT_6 = 0x0020
```




---

* https://www.loc.gov/preservation/digital/formats/fdd/fdd000189.shtml?ref=driverlayer.com/web
* https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bmp.htm
* https://www.digicamsoft.com/bmp/bmp.html
* http://www.kalytta.com/bitmap.h
* http://www.herdsoft.com/ti/davincie/imex3j8i.htm
* http://www.daubnet.com/en/file-format-bmp
* https://en.wikipedia.org/wiki/BMP_file_format
* https://github.com/sol-prog/cpp-bmp-images
* http://paulbourke.net/dataformats/bitmaps/
* http://archive.retro.co.za/CDROMs/DrDobbs/CD%20Release%2012/articles/1995/9503/9503e/9503e.htm
* DR DOBBS
  * https://www.drdobbs.com/architecture-and-design/the-bmp-file-format-part-1/184409517
  * https://www.drdobbs.com/the-bmp-file-format-part-2/184409533?pgno=5
* MICROSOFT
  * https://devblogs.microsoft.com/oldnewthing/20101018-00/?p=12513
  * https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
  * https://docs.microsoft.com/en-us/openspecs/windows_protocols/MS-WMF/792153f4-1e99-4ec8-93cf-d171a5f33903
  * https://github.com/dotnet/wpf/blob/89d172db0b7a192de720c6cfba5e28a1e7d46123/src/Microsoft.DotNet.Wpf/src/WpfGfx/core/glyph/BitmapDbgIO.cpp
* https://www.py4u.net/discuss/128294
* http://www.cplusplus.com/forum/windows/162811/