import wexpect

def run_test():
    
    child = wexpect.spawn('cmd.exe')

    child.sendline('cd D:\\a\\cobra_ui\\cobra_ui\\examples')
    child.expect('>')  # Wait for the command prompt

    # Run the Go program
    child.sendline('go run example3.go')

    child.sendline('25')

    # Wait for the command prompt to confirm the program has finished
    child.expect('>')  

    # Capture the output before the prompt
    output = child.before  
    output_str = output.decode('utf-8') if isinstance(output, bytes) else output  # Decode if it's bytes

    print("Captured output:", output_str)

    # Check if the expected output is in the captured output
    if 'Your entered age is 25' in output_str:
        print("Test passed.")
    else:
        print("Expected output not found.")

    child.close()

if __name__ == "__main__":
    run_test()
