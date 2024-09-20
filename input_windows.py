import wexpect


def run_test():

    child = wexpect.spawn('go run examples/example3.go')

    # Wait for the prompt
    child.expect('Enter your age: ')
    
    # Run the batch file
    child.sendline('25')

    child.send('\n')

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
