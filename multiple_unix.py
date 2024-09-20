import pexpect

def run_test():
  
    child = pexpect.spawn('go run examples/example2.go')

    # Wait for the prompt
    child.expect(pexpect.TIMEOUT, timeout=10)

    child.send('\x1b[B3')  
    child.send('\r')     

   
    child.expect('Selected choice: Python')
  
    child.close()

if __name__ == "__main__":
    run_test()
