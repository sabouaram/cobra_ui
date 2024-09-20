import wexpect

def run_test():
  
    command = r'D:\a\cobra_ui\cobra_ui\examples\example2.exe'
    child = wexpect.spawn(command)
    # Wait for the prompt
    child.expect(pexpect.TIMEOUT, timeout=10)

    child.send('\x1b[B3')  
    child.send('\r')     

   
    child.expect('Selected choice: Python')
  
    child.close()

if __name__ == "__main__":
    run_test()
