import './SearchBar.css';
function SearchBar({ onSearchChange }) {
  return (
    <div className="BarContainer">
      <div className="SearchIcon">
        <svg viewBox="0 0 21 21" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            d="M17.3006 17.3003L13.7973 13.797L17.3006 17.3003ZM13.7973 13.797C14.3596 13.2347 14.8056 12.5671 15.11 11.8324C15.4143 11.0978 15.5709 10.3103 15.5709 9.51514C15.5709 8.71993 15.4143 7.93251 15.11 7.19783C14.8056 6.46315 14.3596 5.79561 13.7973 5.23331C13.235 4.67101 12.5675 4.22497 11.8328 3.92066C11.0981 3.61634 10.3107 3.45972 9.51547 3.45972C8.72026 3.45972 7.93284 3.61634 7.19816 3.92066C6.46349 4.22497 5.79594 4.67101 5.23364 5.23331C4.09803 6.36892 3.46005 7.90914 3.46005 9.51514C3.46005 11.1211 4.09803 12.6614 5.23364 13.797C6.36926 14.9326 7.90948 15.5706 9.51547 15.5706C11.1215 15.5706 12.6617 14.9326 13.7973 13.797V13.797Z"
            stroke="#666666"
            strokeWidth="1.15335"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </div>
      <input
        className="SearchField"
        onChange={(e) => {
          onSearchChange(e.target.value);
        }}
        type="text"
        placeholder="search vendors..."
        spellCheck="false"
      />
    </div>
  );
}

export default SearchBar;
