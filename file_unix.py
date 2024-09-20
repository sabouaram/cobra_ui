import pexpect

def run_test():
  
    child = pexpect.spawn('go run examples/example1.go')

    # Wait for the prompt
    child.expect('Select a file:')

    child.send('/tmp/test_dir/file1.txt')  
    child.send('\r')      # Enter key

   
    child.expect('Selected file full path: /tmp/test_dir/file1.txt')
  
    child.close()

if __name__ == "__main__":
    run_test()
