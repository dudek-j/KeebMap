import './Desktop.css';
import Map from './Map/Map';
import List from './List/List';
import SearchDetails from './SearchDetails/SearchDetails';
import SearchBar from './SearchBar/SearchBar';

function Desktop({
  data,
  highlightItems,
  selectRegion,
  selectedRegions,
  onItemSelected,
  onRegeionSelection,
  onSearchChange,
}) {
  return (
    <div className="App">
      <div className="Header-Container">
        <div className="Header">
          <h3>keebmap.</h3>
        </div>
      </div>
      <div className="Site-Container">
        <div className="Map-Container">
          <Map
            selectedRegions={selectedRegions}
            onRegeionSelection={onRegeionSelection}
          />
        </div>
        <div className="List-Container">
          <SearchBar onSearchChange={onSearchChange} />
          <SearchDetails searchCount={data.length} region={selectRegion} />
          <List
            data={data}
            highlightItems={highlightItems}
            onItemSelected={onItemSelected}
          />
        </div>
      </div>
      <div className="Footer-Container">
        <div className="Footer">
          <p>© 2021 dudek & sjöstrand</p>
          <div className="Blocker" />
          <a href="mailto: contact@keebmap.xyz">contact@keebmap.xyz</a>
        </div>
      </div>
    </div>
  );
}

export default Desktop;
