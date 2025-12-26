from registry import problem_registry as solutions

def main():
    try:
        problem_number = int(input("Enter the problem number: "))
        if problem_number in solutions:
            number = int(input("Enter the upper limit number: "))
            result = solutions[problem_number](number)
            print(f"Result for Problem {problem_number} with limit {number}: {result}")
        else:
            print(f"Problem {problem_number} is not registered.")
    except ValueError:
        print("Invalid input. Please enter numeric values.")

if __name__ == "__main__":
    main()