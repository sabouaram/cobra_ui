; Include necessary constants
#include <Constants.au3>

; Path to your Go program
Local $goProgram = "examples\example3.go"

; Start cmd.exe and run the Go program
Local $cmd = Run(@ComSpec & ' /C "go run ' & $goProgram & '"', "", @SW_HIDE, $STDIN_CHILD + $STDOUT_CHILD)

; Check if the process was created successfully
If $cmd = 0 Then
    ConsoleWrite("Error: Failed to start cmd.exe with Go program." & @CRLF)
    Exit(1)
EndIf

; Wait for the Go program to initialize and prompt for input
Sleep(3000) ; Adjust this delay as necessary

; Send the input '25' followed by ENTER to the Go program
ConsoleWrite("Sending input '25' to the Go program." & @CRLF)
StdinWrite($cmd, "25" & @CRLF)

; Close the STDIN stream if necessary
StdinWrite($cmd)

; Read and capture the output from the Go program
Local $output = ""
While True
    Local $line = StdoutRead($cmd)
    If @error Then ExitLoop ; Exit when there is no more output to read
    $output &= $line
    ConsoleWrite($line) ; Log output for debugging
WEnd

; Check if the output contains the expected result
If StringInStr($output, "Your entered age is 25") Then
    ConsoleWrite("Success: Go program processed input correctly." & @CRLF)
    Exit(0)
Else
    ConsoleWrite("Failure: Go program output did not match expected." & @CRLF)
    Exit(1)
EndIf
