import wexpect
import os

def run_test():
    # Create a temporary batch file
    batch_file_path = 'temp_script.bat'
    with open(batch_file_path, 'w') as f:
        f.write(r'cd D:\\a\\cobra_ui\\cobra_ui\n')
        f.write(r'go run examples/example3.go\n')
    
    # Spawn a new cmd.exe process
    child = wexpect.spawn('cmd.exe')
    
    # Run the batch file
    child.sendline(batch_file_path)

    # Wait for output
    try:
        child.expect('Your entered age is 25')
        print("Test passed.")
    except wexpect.EOF:
        print("EOF reached; command may not have produced expected output.")
    except wexpect.TIMEOUT:
        print("Timeout waiting for expected output.")
    
    # Clean up the temporary batch file
    os.remove(batch_file_path)

    # Close the child process
    child.close()

if __name__ == "__main__":
    run_test()
