# Volkswagen Backend Technical Challenge

The challenge consists of a series of automated robots that are able to receive basic instructions (rotate left, rotate right and move forward) and move across the floor following a grid. 

Since it was not specified how the data ingestion should be for the application, I have decided to use text files that I save in the internal folder and whose path must be passed as a parameter for execution. In the .vscode folder there is a launch.json file that I have decided to keep for simplicity since it is where I specify these paths.

Following what I have done, I have separated the code mainly into 3 parts. The main function, which receives which file to use as input and starts the execution, all the code under the domain folder, which includes the robot, grid and instructions packages together with the service code which is where all the logic of the robots movement is encapsulated and the fileio package which simply has the code that reads the input files and returns the grid, the robots and the instruction lists correctly created.

I have made the assumption that the robots are not on the grid until just before they start moving, i.e., once robot A finishes moving, robot B is placed in its initial position and starts moving. It is important this clarification because I make the checks if a movement is valid in two points. The first one is just before starting to apply the movements to that robot and the second one is just before making a forward movement, to make sure that there are no collisions with other robots and that it does not go out of the defined grid. As the robots move one after another, I only save the final position in a map that I update as the robots finish applying all the instructions that correspond to them.

I have also added an ID to each robot that I am increasing sequentially because I think it can be interesting to know which robot is colliding with one or when reporting errors.

As for the tests that I have applied, I have added a folder called tests that groups all the unit tests of the project and some files with different inputs that slightly vary the initial configuration. The input01.txt file is the default input proposed in the challenge, input02.txt adds two more robots and expands the size of the grid to 10x10 and input03.txt serves to exemplify the errors, both collision and exiting the grid limits.

This is the first time I try to follow a DDD design pattern, at least consciously, but the organization of the code has seemed to me quite intuitive as a whole and I haven't had to change too much my way of working to adapt myself to it, maybe the inclusion of the "service" package is the weirdest thing for me, but I had already created the robot and grid files for example.
