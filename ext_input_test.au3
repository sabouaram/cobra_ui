; Include necessary constants
#include <Constants.au3>

; Path to your Go program
Local $goProgram = "examples\example3.go"
Local $outputFile = @TempDir & "\output.txt"

; Start cmd.exe and run the Go program with output redirection
Local $cmdPath = @ComSpec & ' /C "go run ' & $goProgram & ' > ' & $outputFile & '"'
Local $pid = ShellExecute(@ComSpec, '/C "go run ' & $goProgram & ' > ' & $outputFile & '"', "", "", @SW_SHOW)

; Check if ShellExecute was successful
If $pid = 0 Then
    ConsoleWrite("Error: Failed to start cmd.exe with Go program." & @CRLF)
    Exit(1)
EndIf

; Wait for the Go program to initialize and complete
Sleep(10000) ; Adjust this delay based on the expected execution time

; Read the output from the file
Local $output = FileRead($outputFile)
FileDelete($outputFile) ; Clean up the output file

; Log the output for debugging
ConsoleWrite("Captured output:" & @CRLF & $output & @CRLF)

; Check if the output contains the expected result
If StringInStr($output, "Your entered age is 25") Then
    ConsoleWrite("Success: Go program processed input correctly." & @CRLF)
    Exit(0)
Else
    ConsoleWrite("Failure: Go program output did not match expected." & @CRLF)
    Exit(1)
EndIf
