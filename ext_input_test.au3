; Include constants
#include <Constants.au3>

; Path to the Go example
Local $goProgram = "examples\example3.go"

; Start cmd.exe and run the Go program
Local $cmdPath = @ComSpec & ' /C "go run ' & $goProgram & '"'
Local $cmd = Run($cmdPath, "", @SW_SHOWNORMAL, $STDIN_CHILD + $STDOUT_CHILD + $STDERR_CHILD)

; Check if the process was created successfully
If $cmd = 0 Then
    ConsoleWrite("Error: Failed to start cmd.exe with Go program." & @CRLF)
    Exit(1)
EndIf

; Variable to store the prompt
Local $promptDetected = False

; Loop to continuously check for the prompt
While True
    ; Read from the STDOUT
    Local $line = StdoutRead($cmd)
    If @error Then ExitLoop ; Exit when there is no more output to read

    ; Output the line for debugging
    ConsoleWrite($line)

    ; Check if the prompt is detected
    If StringInStr($line, "Enter your age:") Then
        ; Send the input '25' followed by ENTER to the Go program
        StdinWrite($cmd, "25")
        ConsoleWrite("Sending input '25' to the Go program." & @CRLF)
        Send("{ENTER}")
        $promptDetected = True
    EndIf

    ; Break the loop if the Go program has finished executing
    If Not $promptDetected And StringInStr($line, "Your entered age is 25") Then
        ConsoleWrite("Success: Go program processed input correctly." & @CRLF)
        Exit(0)
    EndIf

    ; Optional
    Sleep(100)
WEnd

; If the loop ends without successful completion
ConsoleWrite("Failure: Go program output did not match expected." & @CRLF)
Exit(1)
