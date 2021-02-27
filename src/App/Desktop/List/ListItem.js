function ListItem({ data, onClick, highlightItem }) {
  return (
    <div
      onClick={() => onClick && onClick(data)}
      className={`ListItem ${highlightItem ? 'HighLightItem' : ''}`}
    >
      {highlightItem ? (
        <div className="HiglightItemDetail">FEATURED</div>
      ) : null}
      <div className="ListItemRow">
        <div className="ListItemWebsite">{data.name.toLowerCase()}</div>
        <div className="ListItemCountry">{data.country.toUpperCase()}</div>
        <div className="ListItemArrow">
          <svg
            width="11"
            height="17"
            viewBox="0 0 11 17"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M2 1.5L9 8.5L2 15.5"
              stroke="#18181D"
              strokeWidth="2.30671"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </div>
      </div>
    </div>
  );
}

export default ListItem;
