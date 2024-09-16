; Run the command prompt with redirection support
Local $pid = Run(@ComSpec & " /c go run examples\example3.go", "", @SW_SHOW, $STDIN_CHILD + $STDOUT_CHILD)

; Check if the process was created successfully
If $pid = 0 Then
    ConsoleWrite("Failed to start the Go program." & @CRLF)
    Exit (1)
EndIf

; Log that the program is running
ConsoleWrite("Running Go Example." & @CRLF)

; Wait for the Go program to start
Sleep(2000)

; Send the input '25' and an ENTER key to the Go program via StdinWrite
StdinWrite($pid, "25" & @CRLF)

ConsoleWrite("Sending argument." & @CRLF)

; Read the output of the Go program
Local $output = ""
While 1
    $line = StdoutRead($pid)
    If @error Then ExitLoop
    $output &= $line
    ConsoleWrite($line) ; Print the output line-by-line for debugging
WEnd

; Split the output into lines
$lines = StringSplit($output, @CRLF)

; Check if there is output and grab the last line
If $lines[0] > 0 Then
    $lastLine = $lines[$lines[0]] ; The last line
    ; Trim any extra spaces
    $lastLine = StringTrimWS($lastLine)

    ; Expected last line
    $expectedLine = "Your entered age is 25"
    
    ; Check if the last line matches the expected line
    If $lastLine = $expectedLine Then
        ConsoleWrite("Success." & @CRLF)
        Exit (0)
    Else
        ConsoleWrite("Fail." & @CRLF)
        Exit(1)
    EndIf
Else
    ConsoleWrite("No output captured." & @CRLF)
    Exit (1)
EndIf
