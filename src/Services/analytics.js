import firebase from 'firebase/app';
import 'firebase/analytics';

const firebaseConfig = {
  apiKey: 'AIzaSyBV88ADUY3QSHelGRwLPvCu-ZWpwmmjK5o',
  authDomain: 'keebmap-3b599.firebaseapp.com',
  projectId: 'keebmap-3b599',
  storageBucket: 'keebmap-3b599.appspot.com',
  messagingSenderId: '384527590644',
  appId: '1:384527590644:web:9298f9a456004393b42aa1',
  measurementId: 'G-99LVGL4WEH',
};

firebase.initializeApp(firebaseConfig);

const Analytics = {
  service: firebase.analytics(),
  logSelectedItem: function (item, selectedRegions) {
    this.service.logEvent(`${item.name}_selected`, {
      selectedRegions: JSON.stringify(selectedRegions),
    });
  },
  logFilterRegionChanged: function (selectedRegions) {
    this.service.logEvent('filter_changed', {
      selectedRegions: JSON.stringify(selectedRegions),
    });
  },
  logSetRegion: function (region) {
    this.service.logEvent(`${region}_region_set`);
  },
  searchQueryTimer: null,
  logSearchQuery: function (query) {
    this.searchQueryTimer && clearTimeout(this.searchQueryTimer);
    this.searchQueryTimer = setTimeout(() => {
      this.service.logEvent('search_query', {
        query: JSON.stringify(query),
      });
    }, 1000);
  },
};

export default Analytics;
