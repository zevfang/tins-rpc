#!/bin/bash

# icon https://www.iconfont.cn/collections/detail?spm=a313x.7781069.1998910419.d9df05512&cid=42825

# use https://github.com/lusingander/fyne-theme-generator to generate theme file

fyne bundle -pkg theme -name ResourceLogoIcon -a -o ../theme/icons.go logo.png

fyne bundle -pkg theme -name ResourcePSquareIcon -a -o ../theme/icons.go p_square.png
fyne bundle -pkg theme -name ResourceSSquareIcon -a -o ../theme/icons.go s_square.png
fyne bundle -pkg theme -name ResourceMSquareIcon -a -o ../theme/icons.go m_square.png

fyne bundle -pkg theme -name ResourceAddIcon -a -o ../theme/icons.go add.png
fyne bundle -pkg theme -name ResourceClearIcon -a -o ../theme/icons.go clear.png
fyne bundle -pkg theme -name ResourceRefreshIcon -a -o ../theme/icons.go refresh.png
fyne bundle -pkg theme -name ResourceRunIcon -a -o ../theme/icons.go run.png