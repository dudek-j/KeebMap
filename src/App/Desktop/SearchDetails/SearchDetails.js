import './SearchDetails.css';
function SearchDetails({ searchCount, region }) {
  return (
    <div className="SearchDetailsContainer">
      <div className="ShowingLabel">{`showing ${searchCount} vendors`}</div>
      <div className="SearchDivider"></div>
      <div className="RegionLabel">
        {(region && region.toUpperCase()) || 'WORLDWIDE'}
      </div>
    </div>
  );
}

export default SearchDetails;
