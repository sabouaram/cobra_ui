import wexpect

def run_test():
    
    child = wexpect.spawn('go run example3.go')

    child.expect('Enter your age: ')

    child.sendline('25')

    child.expect('Your entered age is 25')

    child.close()

if __name__ == "__main__":
    run_test()
