import wexpect


def run_test():
    command = r'cmd.exe /C "go run D:\a\cobra_ui\cobra_ui\examples\example3.go"'
    
    # Spawn the command
    child = wexpect.spawn(command)
    output = child.before
    print("Captured output:", output)

    # Wait for the prompt
    child.expect('Enter your age: ')
    output = child.before
    print("Captured output:", output)
    
    # Run the batch file
    child.sendline('25')

    output = child.before
    print("Captured output:", output)


    # Wait for output
    try:
        child.expect('Your entered age is 25')
        print("Test passed.")
    except wexpect.EOF:
        print("EOF reached; command may not have produced expected output.")
    except wexpect.TIMEOUT:
        print("Timeout waiting for expected output.")
    


    # Close the child process
    child.close()

if __name__ == "__main__":
    run_test()
