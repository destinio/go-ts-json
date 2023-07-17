# 🦸‍♂️ All Things JSON in Go

Are you a JavaScript or TypeScript developer embarking on a grand adventure into the realm of Go, only to find yourself entangled in the treacherous webs of JSON and API calls? Fear not, brave adventurer! With our trusty guide, you shall conquer the challenges and emerge victorious! 🌟

## **⚔️ Making API Calls in Go**

In the land of Go, making API calls may seem unfamiliar, like navigating uncharted territory without your beloved asynchrony and awaits. But fret not, for Go provides a mighty weapon known as the **`net/http`** package, ready to slay APIs with ease! 💪

When wielding this powerful tool, remember an essential rule: always close the response body to avoid resource leaks. Neglecting this duty may lead to a buildup of open network connections and memory leaks, and we wouldn't want that, would we? 🚫💥

To vanquish this threat, you can use the legendary **`defer`** keyword. With a single stroke, it ensures the execution of a function upon the surrounding function's return:

```
goCopy code
resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

```

Behold! The **`defer`** keyword gracefully defers the execution of **`resp.Body.Close()`**, guaranteeing the closure of the response body, even in the face of errors. 🧙‍♂️🚪

By dutifully closing the response body, you not only free precious resources but also demonstrate the mark of a true hero, operating your program with utmost efficiency. 🗝️💡

### **How do we gaze upon the data now? 🤔**

Ah, fear not, dear adventurer! We shall swiftly unveil the secrets hidden within the data and gain insights into our perilous quest.

```
goCopy code
resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
// ...

body, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(body))

```

Alas, this does not satisfy our yearning for more than mere printouts. We must unlock the power of Go's arcane knowledge to progress further.

But before we embark on this journey, let us first understand two essential concepts: 📚✨

- What is a **struct**? 🏰
- The art of encoding and decoding JSON in Go, guided by the mystic **`encoding/json`** package. 🧙‍♂️🌌

## **🏰 Working with Structs in Go**

Structs in the realm of Go may bewilder the minds accustomed to the ways of TypeScript. However, fear not, for they possess great power in handling JSON data. Pay heed to the significance of field tags, for they shall aid you on your quest.

```
goCopy code
// Behold! An example of a struct in Go, adorned with JSON field tags
type Post struct {
    UserID int    `json:"userId"`
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}

```

Structs in Go, much like noble interfaces in TypeScript, shape the form of data and its constituent fields. Yet, in the realm of Go, the fields of a struct possess a steadfast type, unlike the versatile nature of TypeScript interfaces. 🧱💎

## **🧙‍♂️ Encoding and Decoding JSON in Go**

Now, intrepid adventurer, we unveil the mystical art of encoding and decoding JSON in Go. Fear not, for the **`encoding/json`** package shall be your faithful companion on this arcane journey.

Just as JavaScript's **`JSON.parse()`** and **`JSON.stringify()`** reveal the secrets of JSON manipulation, so does Go's **`encoding/json`** package. ✨🔮

### **Let us commence with unmarshalling (decoding)**

Unmarshalling, a wondrous process, transforms the byte sequence bestowed upon us from above into its original data structure, ready for further manipulation.

To wield this power, we shall invoke our Post struct and create an empty list of posts to be "unmarshalled" into. Then, guided by the **`encoding/json`** package, we shall perform our magic. 🪄✨

Finally, the posts shall manifest before us, just as you, dear TypeScript developer, are accustomed to:

```
goCopy code
body, err := io.ReadAll(resp.Body)
if err != nil {
    fmt.Println("error reading body")
    os.Exit(1)
}
posts := []Post{}

// Decode the body into a new list of posts
err = json.Unmarshal(body, &posts)
if err != nil {
    fmt.Println("error unmarshalling", err)
    os.Exit(1)
}

fmt.Println(posts[0].Title)

```

### **But what of marshalling, you ask?**

Ah, marshalling, the art of encoding, allowing us to forge new data to be shared with the realms beyond.

First, we shall create a new post, setting its values using the mighty Post struct.

Next, we invoke the **`json.Marshal`** spell to transform the new post into a variable, ripe for our manipulation.

```
goCopy code
// Prepare for the great marshal!
newPost := Post{
    UserID: 1,
    ID:     1,
    Title:  "New Post",
    Body:   "New Post Body",
}

// Cast the new post into the encoded JSON form
newPostJSON, err := json.Marshal(newPost)
if err != nil {
    fmt.Println("error marshalling", err)
    os.Exit(1)
}

fmt.Println(string(newPostJSON))

```

**And now, let us unleash the POST request to create a new post!**

But beware! To satisfy `http.Post`, we must transform the encoded data into a buffered `bytes.Buffer`, a peculiar requirement known only to the realms of Go. Fear not, for it shall be done!

```
goCopy code
// Prepare for battle, make the POST request!
bufferedPost := bytes.NewBuffer(newPostJSON)

resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bufferedPost)
if err != nil {
    fmt.Println("error POSTing endpoint")
    os.Exit(1)
}
defer resp.Body.Close()

// ...

```

Behold! With a successful request and a triumphant response, we shall deliver the news of victory to the user, for their task is complete!

In the following example, we shall even inspect the expected response status code:

```
goCopy code
// ...
body, err = io.ReadAll(resp.Body)
if err != nil {
    fmt.Println("error reading body")
    os.Exit(1)
}

if resp.StatusCode != 201 {
    fmt.Println("error creating post", resp.StatusCode)
    os.Exit(1)
}

fmt.Println(string(body))

```

### **Headers**

Now, let us discuss the ancient art of setting headers, even crafting custom ones!

First, we shall wield **`http.NewRequest`** instead of **`http.Post`**, granting us the power to customize the request. It bestows upon us a request object, unlike its counterpart, which returns a response.

To make the call and summon the request, we shall employ **`resp, err = http.DefaultClient.Do(req)`**, a union of forces between the request and the mighty **`http.DefaultClient`**.

With the request object in our grasp, we shall invoke **`req.Header.Set`** to bestow upon it the desired headers.

```
goCopy code
// Customize the request headers
bufferedPost := bytes.NewBuffer(newPostJSON)

req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bufferedPost)
if err != nil {
    fmt.Println("error creating request")
    os.Exit(1)
}

req.Header.Set("Content-Type", "application/json")
req.Header.Set("X-My-Header", "my value")

resp, err = http.DefaultClient.Do(req)
if err != nil {
    fmt.Println("error POSTing endpoint")
    os.Exit(1)
}
defer resp.Body.Close()

body, err = io.ReadAll(resp.Body)
if err != nil {
    fmt.Println("error reading body")
    os.Exit(1)
}

if resp.StatusCode != 201 {
    fmt.Println("error creating post", resp.StatusCode)
    os.Exit(1)
}

fmt.Println(string(body))

```

### **Not too shabby, eh?**

Marvelous! Now that we have mastered the ways of working with JSON and embarked on thrilling API requests, let us address the noble art of error handling before we conclude our adventure.

## **🐉 Error Handling**

Ah, the art of handling errors in the realm of Go! Unlike TypeScript, Go requires a more explicit approach, demanding a vigilant eye for potential errors. In TypeScript, the embrace of **`try`** and **`catch`** shields us from harm, but Go demands a different approach.

When making an API call in Go, you must confront the errors that may arise from the HTTP request and response:

```
goCopy code
resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

```

Observe the graceful dance between **`resp`** and **`err`**. The **`http.Get()`** function bequeaths both the HTTP response and an error object. Should the error object not be **`nil`**, it signifies a mishap during the request. In such dire circumstances, we log the error and bid farewell to the program.

Similarly, when decoding JSON in Go, we must confront errors related to the JSON data:

```
goCopy code
var decodedPost Post
err = json.Unmarshal(jsonBytes, &decodedPost)
if err != nil {
    log.Fatal(err)
}

```

Beware the perils that lie within! The **`json.Unmarshal()`** incantation shall summon an error object if it fails to decode the JSON data. Once again, we confront the error explicitly and log it should it arise.

Though it may seem a daunting task, embracing this explicit approach ensures that we catch errors early and handle them with grace.

### **Additional resources and dependencies**

But wait, brave adventurer! Your quest does not end here. There are more treasures to be discovered and dependencies to be explored:

- For further guidance and enlightenment, consult the vast resources available online.
- Delve into the realms of additional dependencies, unveiling the true might of Go in conquering JSON and API calls.

Onward, brave adventurer, for your journey has only just begun! May the code be with you! 🚀

### **Additional Resources:**

1. Go Standard Library - **`net/http`**: **[https://golang.org/pkg/net/http/](https://golang.org/pkg/net/http/)**
2. **`encoding/json`** Package: **[https://golang.org/pkg/encoding/json/](https://golang.org/pkg/encoding/json/)**
3. Go by Example - JSON: **[https://gobyexample.com/json](https://gobyexample.com/json)**
4. JSON and Go Tutorial: **[https://blog.golang.org/json-and-go](https://blog.golang.org/json-and-go)**
5. Go Documentation: **[https://golang.org/doc/](https://golang.org/doc/)**

Happy coding and may your Go adventures be filled with triumph and joy! 🎉✨
