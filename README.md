# GolangUI

## Features
- 對[github.com/zserge/lorca](https://github.com/zserge/lorca)進行修改
- 目標：快速開發一個圖形介面
- 並不是要做package而是做一個模版,之後要簡單做圖形介面只需要複製後修改即可

![](examples/Demo.gif)


## How to use
> 請直接clone,接著按照下面的步驟修改

git clone https://github.com/skynocover/GolangUI

### 基本設定
> 所有的設定都在var內

![](examples/Howtouse.png)

```go
viewtitle  //整個視窗的標題
inputlabel //在輸入框左側的文字
bottonname //於按鈕上顯示的文字
vieweight //視窗的寬度
viewheight //視窗的高度
```

### 程式執行

> 按下按鈕後會做的動作都寫在這裡

![](examples/Func.png)

- 按下按鈕後會將輸入格內的文字變成字串丟入input內
- 直接使用input就可以進行運算
- `document.querySelector('.done').innerText =` 則會將字串輸出到按鈕下面的位置

## Example

[範例](https://github.com/skynocover/Edown)


# Vision

## [1.1.0] 2019-11-20
### Changed
- 修改宣告方式,把ui宣告成全域,讓輸出的方式不只在main()內
> 詳情可以參考[Edown](https://github.com/skynocover/Edown) [1.3.0]版

## [1.0.1] 2019-11-19

### Changed
- 修改宣告方式

----

## [1.0.0] 2019-11-14

最初版本釋出

