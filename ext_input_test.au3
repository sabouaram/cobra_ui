; Open Command Prompt and wait for it to be ready
Run("cmd.exe")
If Not WinWaitActive("[CLASS:ConsoleWindowClass]", "", 10) Then
    ConsoleWrite("Failed to open Command Prompt!" & @CRLF)
    Exit
EndIf

WinActivate("[CLASS:ConsoleWindowClass]")
ConsoleWrite("Command Prompt opened!" & @CRLF)

; Execute the Go program
Send("go run examples\\example3.go{ENTER}")

; Wait for the program to run (Adjust as necessary)
Sleep(5000)

; Send '25' as input
ConsoleWrite("Sending '25' as input..." & @CRLF)
Send("25{ENTER}")

; Wait for the output (Adjust as necessary)
Sleep(5000)

; Capture the entire text from the Command Prompt
$output = WinGetText("[CLASS:ConsoleWindowClass]")
If @error Then
    ConsoleWrite("Failed to capture output from Command Prompt!" & @CRLF)
    Exit
EndIf
ConsoleWrite("Captured output from Command Prompt..." & @CRLF)

; Write captured output for debugging
If Not FileWrite("C:\\temp\\debug_output.txt", $output & @CRLF) Then
    ConsoleWrite("Failed to write to debug_output.txt!" & @CRLF)
    Exit
EndIf
ConsoleWrite("Captured output written to C:\\temp\\debug_output.txt" & @CRLF)

; Split the output into lines and get the last line
$lines = StringSplit($output, @CRLF)
If $lines[0] > 0 Then
    $lastLine = $lines[$lines[0]] ; The last line
    ; Write the last line to a file
    If Not FileWrite("C:\\temp\\output.txt", $lastLine & @CRLF) Then
        ConsoleWrite("Failed to write to output.txt!" & @CRLF)
        Exit
    EndIf
    ConsoleWrite("Last line written to C:\\temp\\output.txt" & @CRLF)
Else
    ConsoleWrite("No output captured!" & @CRLF)
EndIf

Exit
