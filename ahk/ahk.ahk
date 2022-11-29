#p::Pause  ; Win+P
^n::
; run https://www.binance.com/en/convert?fromCoin=BUSD
{
; token icon path
    TokenIconPath := "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook"
; common sleep time
     MouseMoveSpeed := 0
     SleepTimeAfterRefreshPage := 3200
     SleepTimeAfterMouseMove := 50

     SleepTimeForAlternativeClick := 10
     SleepTimeAfterClickActionButton := 1000
     SleepTimeAfterClickInputField := 10
     SleepTimeAfterClickCurrencyEntry := 1300
     SleepTimeAfterClickCurrencyOption := 500

     SleepTimeAfterSendTxt := 10
     SleepTimeAfterSendCurrencyTxt := 600

     SleepAfterClickSlider := 600
; random
    RandomNumber := 0
; login time control
	LoggedInTimeCount := 0
	MaxLoggedInTimeCount := 300 ;18000
	SleepTimeToMeasureLoggedInTime := 100
; Login positions
    MenueX := 907
    MenueY := 144

    LoginX := 774
    LoginY := 197

    ContinueWithGoogleX := 216
    ContinueWithGoogleY := 575

    ChooseAccountX := 891
    ChooseAccountY := 570

; xy_asus_vivobook common positions https://www.binance.com/en/convert?fromCoin=BUSD
	CurrencyGapPixelY := 66

	SelectCurrencyPanX := 951
	SelectCurrencyPanY := 1127

	FromCurrencyValueFieldX := 661
	FromCurrencyValueFieldY := 600
	FromCurrencyEntryX := 868
	FromCurrencyEntryY := 592
	FromCurrencyFieldX := 700
	FromCurrencyFieldY := 664
	FromCurrency1X := 711
	FromCurrency1Y := 740
	FromCurrency2X := 711
	FromCurrency2Y := 831
	FromCurrency3X := 704
	FromCurrency3Y := 920
	FromCurrency4X := 692
	FromCurrency4Y := 1005
	FromCurrency5X := 688
	FromCurrency5Y := 1081

	ToCurrencyValueFieldX := 605
	ToCurrencyValueFieldY := 846
	ToCurrencyEntryX := 864
	ToCurrencyEntryY := 838
	ToCurrencyFieldX := 687
	ToCurrencyFieldY := 668
	ToCurrency1X := 711
	ToCurrency1Y := 745
	ToCurrency2X := 699
	ToCurrency2Y := 828
	ToCurrency3X := 692
	ToCurrency3Y := 916
	ToCurrency4X := 712
	ToCurrency4Y := 997
	ToCurrency5X := 686
	ToCurrency5Y := 1080

	PreviewConversionButtonX := 710
	PreviewConversionButtonY := 940

	PreviewConversionButtonXRisk := 727
	PreviewConversionButtonYRisk := 1029

	ConvertButtonX := 728
	ConvertButtonY := 1095

	ConvertButtonXVolatile := 728
	ConvertButtonYVolatile := 1135

loop
    {
    ; login
        if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\login\login.png")
            {
                MouseMove FoundX+20, FoundY+10, MouseMoveSpeed
                Click
                Sleep 4000
                if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\login\continuewithgoogle.png")
                    {
                       MouseMove FoundX+20, FoundY+10, MouseMoveSpeed
                       Click
                       Sleep 2000
                         if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\login\googlemail.png")
                            {
                                 MouseMove FoundX+20, FoundY+10, MouseMoveSpeed
                                 Click
                                 Sleep 2000
                                 Send "^{Tab}"
                                 Send "{F5}"
                                 Send "^{Tab}"
                                 Sleep 4000
                            }
                    }
             }

        if (RandomNumber = 0)
        {
            MouseMove FromCurrencyEntryX+200, FromCurrencyEntryY+100, MouseMoveSpeed
            RandomNumber := 1
            }
        else
        {
            MouseMove ToCurrencyEntryX+200, ToCurrencyEntryY+100, MouseMoveSpeed
            RandomNumber := 0
            }
        Sleep SleepTimeAfterMouseMove

		; read value data
			FileGainConvertFrom := "gainConvertFrom.txt"
			FileGainConvertTo := "gainConvertTo.txt"
			FileGainConvertValue := "gainConvertValue.txt"

			GainConvertFrom := ""
			GainConvertTo := ""
			GainConvertValue := ""
			if FileExist(FileGainConvertFrom) && FileExist(FileGainConvertTo) && FileExist(FileGainConvertValue) {
				GainConvertFrom := FileRead(FileGainConvertFrom)
				GainConvertTo := FileRead(FileGainConvertTo)
				GainConvertValue := FileRead(FileGainConvertValue)
			}

        if (GainConvertFrom = "T")
            GainConvertTo := "BUSD"

		if (GainConvertFrom != "" && GainConvertTo != "" && GainConvertValue != "")
		{
			; set from token
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeAfterClickInputField
				Send "^a"
				Send GainConvertFrom
				Sleep SleepTimeAfterSendCurrencyTxt
				MouseMove FromCurrency1X, FromCurrency1Y, MouseMoveSpeed

                ; TODO: goto
                ; TODO: not foud abnormal case
				if (GainConvertFrom = "ACA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ACA.png")
                    {        MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterGainConvertFromTokenLocated
                    } else
                            Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "ANT")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ANT.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "AR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else {
                            MouseMove SelectCurrencyPanx, SelectCurrencyPanY, MouseMoveSpeed
                            Click 8
                            Sleep SleepAfterClickSlider
                            if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                            {
                                MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                                Goto LabelAfterGainConvertFromTokenLocated
                            } else
                             Goto LabelGainRefreshPage
                    }
				}
				if (GainConvertFrom = "ATA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ATA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "AUD")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AUD.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "AVA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AVA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "BAR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BAR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "BNB")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BNB.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "BTC")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BTC.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "COS")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\COS.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "ETH")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ETH.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "FOR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\FOR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "GAL")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GAL.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "GBP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GBP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "GLM")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GLM.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "OG")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OG.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "OM")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else {
                           MouseMove SelectCurrencyPanX, SelectCurrencyPanY, MouseMoveSpeed
                           Click 4
                           Sleep SleepAfterClickSlider
                           if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                            {
                              MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                              Goto LabelAfterGainConvertFromTokenLocated
                             } else
                              Goto LabelGainRefreshPage
                    }
				}
				if (GainConvertFrom = "ONT")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ONT.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "OP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "ORN")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ORN.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "PHA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\PHA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "REP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\REP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "SC")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\SC.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "T")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else {
                            MouseMove SelectCurrencyPanX, SelectCurrencyPanY, MouseMoveSpeed
                            LoopTs := 0
                            LoopT:
                            Click 8
                            Sleep SleepAfterClickSlider
                            if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                             {
                                 MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                                  Goto LabelAfterGainConvertFromTokenLocated
                              } else {
                                LoopTs += 1
                                if LoopTs > 20
                                    Goto LabelGainRefreshPage
                                else
                                    Goto LoopT
                              }
                    }
				}
				if (GainConvertFrom = "WIN")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\WIN.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}
				if (GainConvertFrom = "YFI")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\YFI.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterGainConvertFromTokenLocated
                    } else
                             Goto LabelGainRefreshPage
				}


                LabelAfterGainConvertFromTokenLocated:
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token: BUSD, or USDT
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeAfterClickInputField
				Send "^a"
				Send GainConvertTo
				Sleep SleepTimeAfterSendCurrencyTxt
				MouseMove ToCurrency1X, ToCurrency1Y, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeAfterClickInputField
				Send GainConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
			LabelGainRefreshPage:
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage ;todo: window wait
				Send "^{Tab}"
		}

		FileLossConvertFrom := "lossConvertFrom.txt"
		FileLossConvertTo := "lossConvertTo.txt"
		FileLossConvertValue := "lossConvertValue.txt"

		LossConvertFrom := ""
		LossConvertTo := ""
		LossConvertValue := ""
		if FileExist(FileLossConvertFrom) && FileExist(FileLossConvertTo) && FileExist(FileLossConvertValue) {
			LossConvertFrom := FileRead(FileLossConvertFrom)
			LossConvertTo := FileRead(FileLossConvertTo)
			LossConvertValue := FileRead(FileLossConvertValue)
		}

     if (LossConvertTo = "T")
            LossConvertFrom := "BUSD"

		if (LossConvertFrom != "" && LossConvertTo != "" && LossConvertValue != "")
		{
			; set from tokenMOVR, limit to: BUSD, USDT
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send "^a"
				Send LossConvertFrom
				Sleep SleepTimeAfterSendCurrencyTxt
				MouseMove FromCurrency1X, FromCurrency1Y, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeAfterClickInputField
				Send "^a"
				Send LossConvertTo
				Sleep SleepTimeAfterSendCurrencyTxt
				MouseMove ToCurrency1X, ToCurrency1Y, MouseMoveSpeed
				if (LossConvertTo = "ACA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ACA.png")
                    {        MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                            Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "ANT")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ANT.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "AR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else {
                            MouseMove SelectCurrencyPanX, SelectCurrencyPanY, MouseMoveSpeed
                            Click 8
                            Sleep SleepAfterClickSlider
                            if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AR.png")
                            {
                                    MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                                    Goto LabelAfterLossConvertToTokenLocated
                            } else
                                Goto LabelLossRefreshPage
                    }
				}
				if (LossConvertTo = "ATA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ATA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "AUD")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AUD.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "AVA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\AVA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "BAR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BAR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "BNB")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BNB.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "BTC")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\BTC.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "COS")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\COS.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "ETH")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ETH.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "FOR")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\FOR.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "GAL")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GAL.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "GBP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GBP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "GLM")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\GLM.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "OG")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OG.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "OM")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else {
                            MouseMove SelectCurrencyPanX, SelectCurrencyPanY, MouseMoveSpeed
                            Click 4
                            Sleep SleepAfterClickSlider
                            if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OM.png")
                             {
                                  MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                                  Goto LabelAfterLossConvertToTokenLocated
                              } else
                                Goto LabelLossRefreshPage
                    }
				}
				if (LossConvertTo = "ONT")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ONT.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "OP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\OP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "ORN")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\ORN.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "PHA")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\PHA.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "REP")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\REP.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "SC")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\SC.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
		        if (LossConvertTo = "T")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                           Goto LabelAfterLossConvertToTokenLocated
                    } else {
                            MouseMove SelectCurrencyPanX, SelectCurrencyPanY, MouseMoveSpeed
                            LossLoopTs := 0
                            LossLoopT:
                            Click 8
                            Sleep SleepAfterClickSlider
                            if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\T.png")
                             {
                                 MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                                  Goto LabelAfterLossConvertToTokenLocated
                              } else {
                                LossLoopTs += 1
                                if LossLoopTs > 20
                                    Goto LabelLossRefreshPage
                                else
                                    Goto LossLoopT
                              }
                    }
				}
				if (LossConvertTo = "WIN")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\WIN.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}
				if (LossConvertTo = "YFI")
				{
				    MouseMove FromCurrencyEntryX, FromCurrencyEntryY, MouseMoveSpeed
				    if ImageSearch(&FoundX, &FoundY,0, 0, A_ScreenWidth, A_ScreenHeight, "W:\github.com\ljg-cqu\binance\biance\static\tokenicon\asusvivobook\YFI.png")
                    {
                            MouseMove FoundX+20, FoundY+20, MouseMoveSpeed
                            Goto LabelAfterLossConvertToTokenLocated
                    } else
                             Goto LabelLossRefreshPage
				}

				LabelAfterLossConvertToTokenLocated:
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeAfterClickInputField
				Send LossConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				;Sleep SleepTimeForAlternativeClick
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile, MouseMoveSpeed
				;Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
			LabelLossRefreshPage:
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage
				Send "^{Tab}"
		}
    }
}