^8::
{
;CoordMode "Pixel"
;CoordMode, Pix8el, Window

MouseMove 583,475
Click
Sleep 1500
MouseMove 471,429
Click
Send "AVA"
Sleep 600
if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\E52660\AVA.png")
       {
        MsgBox "The icon was found at " FoundX "x" FoundY
        MouseMove FoundX+20, FoundY+20
        Click
        }
    else
        MsgBox "Icon could not be found on the screen."
}