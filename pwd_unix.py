import pexpect

def run_test():
  
    child = pexpect.spawn('go run examples/example4.go')

    # Wait for the prompt
    child.expect('Enter your password: ')

    child.send('123456')  
    child.send('\r')      # Enter key

   
    child.expect('Password entered => 123456')
  
    child.close()

if __name__ == "__main__":
    run_test()
