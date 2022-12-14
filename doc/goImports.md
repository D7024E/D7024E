# Go imports
All local packages exist on combined virtual and real filepath, the virtual path is defined in go.mod and is the prefix for all real filepaths.
For instance:<br />
main<br />
----main.go                 <br />
----second.go               <br />
----go.mod                  <br />
----article/                <br />
-----------article.go       <br />
-----------review.go        <br />
----logic/                  <br />
---------logic.go           <br />

In this filestructure the virtual path is defined in go.mod, however main.go will still exist in the package "main" as it is required as the base. Lets say that the virtual path is defined as "D7024E" and we want to import the package where article exists from the "logic.go" file and we know that it happens to be called "article". We then call the import for the virtual path followed by the real path. Note that the real path always starts from the "main" package:

import (<br />
&nbsp;&nbsp;&nbsp;&nbsp;"D7024E/article"<br />
)<br />

We can now use the public namespace from the article.go file, but we can also use the public namespace from the review.go file. This is because they both must exist in the same package, as they exist in the same folder and all files in the same folder must be in the same package by definition in go. Also note that we do not need to import the specific package by name, it is enough to import the folder where the package exists since all files in that folder must exist in the same package. The package also does not have to use the same name as its folder, although it is probably best to use the folder name to reduce confusion.

Now lets say that we already use the article packages name locally in the logic.go file for something else. No problem we can set up an alias for it in our import, simply prepend the import string with the alias you want to use for it. For example if we want to use the alias "fakeNews":

import (<br />
&nbsp;&nbsp;&nbsp;&nbsp;fakeNews "D7024E/article"<br />
)<br />

We now import the article package under the name "fakeNews", simple as that.

# Tips and tricks
If you want go to stop complaining that you import something but don't use it, you can set a wildcard alias for the package. This is done by setting a underscore, "_", as the package alias. For example, if we want to import the article folders package but wont use it yet: <br />
import (<br />
&nbsp;&nbsp;&nbsp;&nbsp; _ "D7024E/article"<br />
)<br />