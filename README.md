# tutorial-minimal-example

- Note: 執行指令可按指令左邊的播放鍵執行。

## 觀察檔案
對下面兩個檔案進行觀察
- `minimal.bpf.c` :  
  這個檔案就是我們要載入Kernel的eBPF程式

- `main.go` :  
  這個檔案就是我們在User-space的前端，負責載入eBPF程式至Kernel中。

## 執行eBPF程式

```bash
make run
```

## 察看Kernel中的trace訊息

觀察trace中是否有印出我們需要的資訊，i.e. 呼叫 `write` syscall 的 PID 。

```bash
make trace
```

## 修改eBPF程式

試著修改 `minimal.bpf.c` 中印出訊息的地方，並重複執行、觀察trace的流程。看看顯示的訊息是否會改變。
