import wexpect

def run_test():
    
    child = wexpect.spawn('go run example3.go')

    child.sendline('25')

    # Capture the output before the prompt
    output = child.before  

    print(output)

    child.close()

if __name__ == "__main__":
    run_test()
