def print_menu():
    menu_options = {
        1: 'Option 1',
        2: 'Option 2',
        3: 'Option 3',
        4: 'Option 4',
        5: 'Exit'
    }

    for key in menu_options.keys():
        print(key, '--', menu_options[key])

    option = int(input('\nEnter your choice: '))

    if option == 1:
        print('\nHandle option \'Option 1\'\n')
    elif option == 2:
        print('\nHandle option \'Option 2\'\n')
    elif option == 3:
        print('\nHandle option \'Option 3\'\n')
    elif option == 4:
        print('\nHandle option \'Option 4\'\n')
    elif option == 5:
        print('\nCatch you later.')
        exit()
    else:
        print('\nInvalid option. Please enter a valid number.\n')
