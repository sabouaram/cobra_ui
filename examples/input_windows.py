import wexpect

def run_test():

    command = r'D:\a\cobra_ui\cobra_ui\examples\example3.exe'
    
    child = wexpect.spawn(command)

    child.expect('Enter your age: ')

    print(child.before)

    child.send('25\n\r')

    print(child.before)

    child.expect('Your entered age is 25')

    print(child.before)

    child.close()

if __name__ == "__main__":
    run_test()
