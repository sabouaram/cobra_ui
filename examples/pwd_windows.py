import wexpect

def run_test():
  
    command = r'D:\a\cobra_ui\cobra_ui\examples\example4.exe'
    child = wexpect.spawn(command)

    # Wait for the prompt
    child.expect('Enter your password: ')

    child.send('123456')  
    child.send('\r')      # Enter key
  
    child.close()

if __name__ == "__main__":
    run_test()
