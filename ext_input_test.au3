; Run cmd.exe and execute 'go run' in it
Local $cmd = Run("go run examples\example3.go", "", @SW_SHOW, $STDIN_CHILD + $STDOUT_CHILD + $STDERR_CHILD)

; Check if the process was created successfully
If $cmd = 0 Then
    ConsoleWrite("Failed to start cmd.exe with Go program." & @CRLF)
    Exit(1)
EndIf

ConsoleWrite("Go example started successfully." & @CRLF)
Sleep(3000) ; Wait for Go program to initialize

; Send the input '25' followed by ENTER to the Go program
StdinWrite($cmd, "25" & @CRLF)
ConsoleWrite("Sending argument '25'." & @CRLF)

; Non-blocking loop to read stdout and stderr
Local $output = ""
While ProcessExists($cmd)
    ; Read from stdout
    Local $stdout = StdoutRead($cmd)
    If @error Then ExitLoop
    If $stdout <> "" Then
        $output &= $stdout
        ConsoleWrite($stdout) ; Log stdout for debugging
    EndIf

    ; Read from stderr
    Local $stderr = StderrRead($cmd)
    If @error Then ExitLoop
    If $stderr <> "" Then
        $output &= $stderr
        ConsoleWrite("Error: " & $stderr) ; Log stderr for debugging
    EndIf

    ; Add a short sleep to prevent high CPU usage
    Sleep(100)
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
