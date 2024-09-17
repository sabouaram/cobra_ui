; Run cmd.exe and execute 'go run' in it
Local $cmd = Run(@ComSpec & " /c go run examples\example3.go", "", @SW_SHOW, $STDIN_CHILD + $STDOUT_CHILD + $STDERR_CHILD)

; Check if the process was created successfully
If $cmd = 0 Then
    ConsoleWrite("Failed to start cmd.exe with Go program." & @CRLF)
    Exit(1)
EndIf


ConsoleWrite("Go example started successfully." & @CRLF)

; Wait for the Go program to initialize
Sleep(3000) ; Adjust if necessary

; Send the input '25' followed by ENTER to the Go program
StdinWrite($cmd, "25" & @CRLF)
ConsoleWrite("Sending argument '25'." & @CRLF)

; Read and capture the output from the Go program
Local $output = ""
While 1
    Local $line = StdoutRead($cmd)
    If @error Then ExitLoop
    $output &= $line
    ConsoleWrite($line) ; Log output for debugging
WEnd

; Close the STDIN stream if necessary
StdinWrite($cmd)

; Check if the output contains the expected result
If StringInStr($output, "Your entered age is 25") Then
    ConsoleWrite("Success." & @CRLF)
    Exit(0)
Else
    ConsoleWrite("Fail." & @CRLF)
    Exit(1)
EndIf
