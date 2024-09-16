; Start cmd.exe
Local $cmd = Run("cmd.exe", "", @SW_SHOW, $STDIN_CHILD + $STDOUT_CHILD)

ConsoleWrite("here." & @CRLF)



; Activate the command prompt window
WinActivate("[CLASS:ConsoleWindowClass]")

ConsoleWrite("Running Go Example." & @CRLF)

; Run the Go program via 'go run' inside the command prompt
StdinWrite($cmd, "go run examples\example3.go" & @CRLF)
ConsoleWrite("Example Running." & @CRLF)

; Wait for the Go program to initialize (adjust sleep if necessary)
Sleep(3000)

; Send the input '25' followed by ENTER to the Go program
StdinWrite($cmd, "25" & @CRLF)
ConsoleWrite("Sending argument." & @CRLF)

; Capture the program's output
Local $output = ""
While 1
    Local $line = StdoutRead($cmd)
    If @error Then ExitLoop
    $output &= $line
    ConsoleWrite($line) ; Log the output for debugging
WEnd

; Close stdin to signal end of input if needed
StdinWrite($cmd)

; Check if the output matches the expected result
If StringInStr($output, "Your entered age is 25") Then
    ConsoleWrite("Success." & @CRLF)
    Exit(0)
Else
    ConsoleWrite("Fail." & @CRLF)
    Exit(1)
EndIf
