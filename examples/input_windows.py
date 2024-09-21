import wexpect

def run_test():

    command = r'D:\a\cobra_ui\cobra_ui\examples\example3.exe'

    
    
    child = wexpect.spawn(command)

    child.expect('Enter your age: ',timeout=5)


    child.send('25\n\r')

    child.expect('')

    with open('D:\\a\\cobra_ui\\cobra_ui\\examples\\output.log', 'w') as log_file:
        # Write the captured output to the file
        log_file.write(child.before)
    
    

    child.close()

if __name__ == "__main__":
    run_test()
