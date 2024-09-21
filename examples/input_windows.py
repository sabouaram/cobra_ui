import wexpect

def run_test():

    command = r'D:\a\cobra_ui\cobra_ui\examples\example3.exe'
    
    child = wexpect.spawn(command + ' > output.log')

    child.expect('Enter your age: ',timeout=5)

    print(child.before)

    child.send('25\n\r')

    child.close()

if __name__ == "__main__":
    run_test()
