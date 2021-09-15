import URLShortener from './URLShortener';
import './App.css'

const App = () => {
  return <URLShortener 
    containerStyle= { {  } }
    formStyle=      { {  } }
    buttonStyle=    { { border: "solid 2px black", borderRadius: "3px", padding: "5px" } }
    inputStyle=     { {color: "rgb(225, 225, 225)", border: "none", background: "rgb(146, 147, 148)", outline: "solid 2px grey",
                      padding: "8px", borderRadius: "5px", fontWeight: "600", caretShape: "block", textShadow: "1px 1px 2px black"}}
    errorStyle=     { { color: "#fff", background: "rgb(227, 32, 32)", padding: "10px", borderRadius: "5px", fontWeight: "600"} }
    linkStyle=      { { color: "#000" } }
  />

}
export default App;
