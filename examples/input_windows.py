import wexpect

def run_test():

    command = r'D:\a\cobra_ui\cobra_ui\examples\example3.exe'

    
    
    child = wexpect.spawn(command)

    child.expect('Enter your age: ',timeout=5)


    child.send('25\n\r')

    child.expect('')
    print(child.before)

    child.close()

if __name__ == "__main__":
    run_test()
