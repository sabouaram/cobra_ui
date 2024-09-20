import wexpect

def run_test():
    
    child = wexpect.spawn('cmd.exe')
    child.sendline('cd  D:\\a\\cobra_ui\\cobra_ui')
    child.sendline('go run examples\\example3.go')
 

    # Wait for the prompt
    child.expect('Enter your age: ')

    child.sendline('25')      # Enter key

   
    child.expect('Your entered age is 25')
  
    child.close()

if __name__ == "__main__":
    run_test()
