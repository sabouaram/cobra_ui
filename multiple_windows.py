import wexpect

def run_test():
  
    child = wexpect.spawn('go run examples/example1.go')

    # Wait for the prompt
    child.expect('Select a file:')

    child.send('C:\\temp\\test_dir\\file1.txt')  
    child.send('\r')      # Enter key

   
    child.expect('Selected file: file1.txt')
  
    child.close()

if __name__ == "__main__":
    run_test()
