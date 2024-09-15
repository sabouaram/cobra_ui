; Open Command Prompt and wait for it to be ready
Run("cmd.exe")
If Not WinWaitActive("[CLASS:ConsoleWindowClass]", "", 10) Then
    Exit 1 ; Exit with code 1 to indicate failure
EndIf

WinActivate("[CLASS:ConsoleWindowClass]")

; Execute the Go example
Send("go run examples\example3.go{ENTER}")

; Wait for the program to run (Adjust as necessary)
Sleep(5000)

; Send '25' as input
Send("25{ENTER}")

; Wait for the output (Adjust as necessary)
Sleep(5000)

; Capture the entire text from the Command Prompt
$output = WinGetText("[CLASS:ConsoleWindowClass]")
If @error Then
    Exit 1 ; Exit with code 1 to indicate failure
EndIf

; Split the output into lines and get the last line
$lines = StringSplit($output, @CRLF)
If $lines[0] > 0 Then
    $lastLine = $lines[$lines[0]] ; The last line
    ; Trim any extra spaces
    $lastLine = StringTrimWS($lastLine)
    
    ; Expected last line
    $expectedLine = "Your entered age is 25"
    
    ; Check if the last line matches the expected line
    If $lastLine = $expectedLine Then
        Exit 0 ; Exit with code 0 to indicate success
    Else
        Exit 1 ; Exit with code 1 to indicate failure
    EndIf
Else
    Exit 1 ; Exit with code 1 to indicate failure
EndIf
