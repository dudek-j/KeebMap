import React, { useState } from 'react';
import Desktop from './Desktop';
import Mobile from './Mobile';
import Analytics from '../Services/analytics';
import { isMobile } from 'react-device-detect';

import { openInNewTab, allValuesFalse } from './Utility';

function App({ data }) {
  const [searchQuery, setSearchQuery] = useState('');
  const [mapState, setMapState] = useState({
    europe: false,
    uk: false,
    canada: false,
    us: false,
    oceania: false,
    asia: false,
    china: false,
    'latin america': false,
    africa: false,
  });

  data = data.filter((i) => !i.hide);

  //Filter region Analytics
  // useEffect(() => {
  //   if (!allValuesFalse(selectedRegions)) {
  //     Analytics.logFilterRegionChanged(selectedRegions);
  //   }
  // }, [selectedRegions]);

  //Search Query Analytics
  // useEffect(() => {
  //   if (searchQuery.length !== 0) {
  //     Analytics.logSearchQuery(searchQuery);
  //   }
  // }, [searchQuery]);

  function filterData(regions, query) {
    const allRegionsFalse = allValuesFalse(regions);

    const searched = data.filter((vendor) =>
      vendor.name.toLowerCase().includes(query.toLowerCase())
    );

    if (allRegionsFalse) {
      return searched;
    } else {
      const filtered = searched.filter((vendor) => regions[vendor.region]);
      return filtered;
    }
  }

  function onListItemSelcted(item) {
    Analytics.logSelectedItem(item, mapState);
    openInNewTab(item.url);
  }

  function onRegionSelected(region) {
    const copy = { ...mapState };

    //Set all not pressed to false
    Object.keys(copy).forEach((v) => {
      if (v !== region) {
        copy[v] = false;
      }
    });

    //Toggle the pressed region
    const val = !mapState[region];

    copy[region] = val;
    setMapState(copy);
    val && Analytics.logSetRegion(region);
  }

  const processedData = filterData(mapState, searchQuery);
  const selectedRegion = Object.entries(mapState).find((e) => e[1]);

  return isMobile ? (
    <Mobile />
  ) : (
    <Desktop
      data={processedData}
      highlightItems={!allValuesFalse(mapState)}
      onItemSelected={onListItemSelcted}
      mapRegions={mapState}
      onRegionSelected={onRegionSelected}
      onSearchChange={setSearchQuery}
      selectRegion={selectedRegion && selectedRegion[0]}
    />
  );
}

export default App;
