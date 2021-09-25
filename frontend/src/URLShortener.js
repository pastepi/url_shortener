import { useState } from "react"


const defaultContainerStyle = {minHeight: "150px", margin: "0", minWidth: "300px", display: "flex", alignItems: "center", justifyContent: "center"};
const defaultFormStyle = {display: "flex", flexDirection: "column", alignItems: "center", minHeight: "150px", justifyContent: "space-around"};
const defaultButtonStyle = {fontSize: "1rem"};
const defaultInputStyle = {fontSize: "1rem"};
const defaultLinkStyle = {fontSize: "1.2rem"};
const defaultErrorStyle = {color: "red"};


const myHost = process.env.REACT_APP_BACKEND_URL;


const URLShortener = ({ containerStyle, formStyle, buttonStyle, inputStyle, linkStyle, errorStyle }) => {
    
    const [url, setURL] = useState("");
    const [shortURL, setShortURL] = useState(null);
    const [error, setErrorState] = useState(false);
  
    const handleURLChange = (e) => {
      setURL(e.target.value)
    }
  
    const isValidURL = (srcURL) => {
      let pattern = new RegExp('^(https?:\\/\\/)?'+         // protocol
        '((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|'+ // domain name
        '((\\d{1,3}\\.){3}\\d{1,3}))'+                      // OR ip (v4) address
        '(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*'+                  // port and path
        '(\\?[;&a-z\\d%_.~+=-]*)?'+                         // query string
        '(\\#[-a-z\\d_]*)?$','i');                          // fragment locator
    
        return pattern.test(srcURL);
    }

    const handleButtonClick = (e) => {
      e.preventDefault()
      if (error) {
        setErrorState(false)
      }

      if (isValidURL(url)) {
        let srcURL = url;

        // Checks for the scheme/protocol of the link 
        // - if no valid one exists, adds "HTTPS"
        try {
          new URL(srcURL);
        } catch (_) { // MalformedURLException
          srcURL = "https://" + srcURL;
          setURL(srcURL);
        }

        fetch(myHost + "/URL", {
          method: "POST", 
          headers: {
          "Content-Type": "application/json"
          },
          body: JSON.stringify({url:srcURL})
      }).then(response => response.json()
      ).then(data => setShortURL(data.ShortURL))
      } else {
        setErrorState(true);
        setShortURL(null);
      }
    }

    const styleElements = (defStyle, style) => {
      return style ? {...defStyle, ...style} : defStyle
    }
    
    return (
      <div style={ styleElements(defaultContainerStyle, containerStyle) }>
        <form style={ styleElements(defaultFormStyle, formStyle) }>
          <input          style={ styleElements(defaultInputStyle, inputStyle) } type="text" value={url} onChange={handleURLChange} spellCheck="false" placeholder="Your URL"/>
          <button         style={ styleElements(defaultButtonStyle, buttonStyle) } onClick={handleButtonClick}>Shorten URL</button>
          {shortURL && <a style={ styleElements(defaultLinkStyle, linkStyle) } href={`${myHost}/${shortURL}`} rel="noreferrer" target="_blank" >{`${myHost}/${shortURL}`}</a>}
          {error && <div  style={ styleElements(defaultErrorStyle, errorStyle) }>URL is not valid</div>}
        </form>
      </div>
    );
}

export default URLShortener;