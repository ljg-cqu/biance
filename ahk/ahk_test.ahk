^8::
{
CoordMode "Pixel"

if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\AR.png")
        MsgBox "The icon was found at " FoundX "x" FoundY
        Sleep 3000
        MouseMove FoundX, FoundY

    else
        MsgBox "Icon could not be found on the screen."
}