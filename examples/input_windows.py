import wexpect

def run_test():

    command = r'D:\\a\cobra_ui\cobra_ui\examples\example3\example3.exe'
    
    child = wexpect.spawn(command)

    child.expect('Enter your age: ')

    child.sendline('25')

    child.expect('Your entered age is 25')

    child.close()

if __name__ == "__main__":
    run_test()
