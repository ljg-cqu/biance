^n::
; run https://www.binance.com/en/convert?fromCoin=BUSD
{
; refresh page
	PageIdleNumber := 0
	MaxPageIdleNumber := 60
	SleepTimeToAdjustIdleNumber := 1000
; common sleep time
    SleepTimeAfterRefreshPage := 5000
	SleepTimeAfterMouseMove := 1

	SleepTimeForAlternativeClick := 1
	SleepTimeAfterClickActionButton := 500
	SleepTimeAfterClickInputField := 1
    SleepTimeAfterClickCurrencyEntry := 1500
    SleepTimeAfterClickCurrencyOption := 500

    SleepTimeAfterSendTxt := 1

; common positions
	CurrencyGapPixelY := 66

	FromCurrencyValueFieldX := 428
	FromCurrencyValueFieldY := 478
	FromCurrencyEntryX := 577
	FromCurrencyEntryY := 472
	FromCurrencyFieldX := 460
	FromCurrencyFieldY := 425
	FromCurrency1X := 462
	FromCurrency1Y := 486
	FromCurrency2X := 458
	FromCurrency2Y := 552
	FromCurrency3X := 467
	FromCurrency3Y := 617
	FromCurrency4X := 470
	FromCurrency4Y := 690
	FromCurrency5X := 457
	FromCurrency5Y := 760

	ToCurrencyValueFieldX := 400
	ToCurrencyValueFieldY := 678
	ToCurrencyEntryX := 586
	ToCurrencyEntryY := 668
	ToCurrencyFieldX := 468
	ToCurrencyFieldY := 425
	ToCurrency1X := 467
	ToCurrency1Y := 486
	ToCurrency2X := 467
	ToCurrency2Y := 549
	ToCurrency3X := 464
	ToCurrency3Y := 626
	ToCurrency4X := 454
	ToCurrency4Y := 690
	ToCurrency5X := 452
	ToCurrency5Y := 756

	PreviewConversionButtonX := 471
	PreviewConversionButtonY := 753

	PreviewConversionButtonXRisk := 479
	PreviewConversionButtonYRisk := 839

	ConvertButtonX := 464
	ConvertButtonY := 768

	ConvertButtonXVolatile := 474
	ConvertButtonYVolatile := 808

loop
    {
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

		if (GainConvertFrom != "" && GainConvertTo != "" && GainConvertValue != "")
		{
			; set from token
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertFrom
				Sleep SleepTimeAfterSendTxt
				MouseMove FromCurrency1X, FromCurrency1Y, 200

				; AR, T, AST
				if (GainConvertFrom = "ATA") ; TODO: add more token check ; TODO: reuse, image s.
				{
				MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "PHA")
				{
				MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "WIN")
				{
				MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "OP")
				{
				MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "GLM")
				{
				MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "AUD") ; TODO: add more token check
				{
					MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "BAR")
				{
					MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "COS") ;
				{
					MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "ETH") ;
				{
					MouseMove FromCurrency2X, FromCurrency2Y
				}
				if (GainConvertFrom = "AVA")
				{
					MouseMove FromCurrency3X, FromCurrency3Y
				}
				if (GainConvertFrom = "OM")
				{
					MouseMove FromCurrency5X, FromCurrency5Y
				}

				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertTo
				Sleep SleepTimeAfterSendTxt
				MouseMove ToCurrency1X, ToCurrency1Y
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send GainConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY
				Sleep SleepTimeAfterMouseMove
				Click
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage ;todo: window wait
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

		if (LossConvertFrom != "" && LossConvertTo != "" && LossConvertValue != "")
		{
			; set from tokenMOVR, limit to: BUSD, USDT
				MouseMove FromCurrencyEntryX, FromCurrencyEntryY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove FromCurrencyFieldX, FromCurrencyFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertFrom
				Sleep SleepTimeAfterSendTxt
				MouseMove FromCurrency1X, FromCurrency1Y
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyOption

			; set to token
				MouseMove ToCurrencyEntryX, ToCurrencyEntryY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry
				MouseMove ToCurrencyFieldX, ToCurrencyFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertTo
				Sleep SleepTimeAfterSendTxt
				MouseMove ToCurrency1X, ToCurrency1Y
				; OM, AST, T
				if (LossConvertTo = "AR") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "ONT") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "OG") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "GAL") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "PHA") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "BAR") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "COS") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				if (LossConvertTo = "OM") ;
				{
					MouseMove ToCurrency3X, ToCurrency3Y
				}
				if (LossConvertTo = "AVA") ;
				{
					MouseMove ToCurrency2X, ToCurrency2Y
				}
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickCurrencyEntry

			; set convert value
				MouseMove FromCurrencyValueFieldX, FromCurrencyValueFieldY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickInputField
				Send LossConvertValue
				Sleep SleepTimeAfterSendTxt

			; preview conversion
				MouseMove PreviewConversionButtonX, PreviewConversionButtonY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove PreviewConversionButtonXRisk, PreviewConversionButtonYRisk
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; convert
				MouseMove ConvertButtonX, ConvertButtonY
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeForAlternativeClick
				MouseMove ConvertButtonXVolatile, ConvertButtonYVolatile
				Sleep SleepTimeAfterMouseMove
				Click
				Sleep SleepTimeAfterClickActionButton

			; refresh web page
				Send "{F5}" ; TODO: remove it after balance check by server side
				Sleep SleepTimeAfterRefreshPage
		}

		if (GainConvertFrom != "" && GainConvertTo != "" && GainConvertValue != "" && LossConvertFrom != "" && LossConvertTo != "" && LossConvertValue != "")
		{
			PageIdleNumber += 1
			Sleep SleepTimeToAdjustIdleNumber
		} else
		{
			PageIdleNumber := 0
		}

		if (PageIdleNumber > MaxPageIdleNumber )
		{
			Send "{F5}"
			Sleep SleepTimeAfterRefreshPage

		; TODO: extract function
	    }
    }
}