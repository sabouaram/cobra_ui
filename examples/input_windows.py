import wexpect

def run_test():
    child = wexpect.spawn('go run example3.go')
  
    child.logfile = sys.stdout  # Log output for debugging
  
    output = child.before  

    print("Captured output:", output)


    # Wait for the prompt
    child.expect('Enter your age: ')
    output = child.before
    print("Captured output:", output)
    
    # Run the input
    child.sendline('25')


    try:
        child.expect('Your entered age is 25', timeout=60)  # Adjust timeout
        print("Test passed.")
    except wexpect.EOF:
        print("EOF reached; command may not have produced expected output.")
    except wexpect.TIMEOUT:
        print("Timeout waiting for expected output.")
        output = child.before.decode()  # Capture output before timeout
        print("Captured output before timeout:", output)

    child.close()

if __name__ == "__main__":
    run_test()

