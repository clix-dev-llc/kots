import React from "react";
import classNames from "classnames";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";

import { isHelmChart } from "@src/utilities/utilities";
import subNavConfig from "@src/config-ui/subNavConfig";

export default function SubNavBar(props) {
  const { className, activeTab, watch, isVeleroInstalled, isAccess } = props;
  let { slug } = watch;

  if (isHelmChart(watch)) {
    slug = `helm/${watch.id}`;
  }
  const kotsSequence = watch.currentSequence;

  const accessConfig = [
    {
      tabName: "configure-ingress",
      displayName: "Configure Ingress",
      to: () => `/access/configure-ingress`,
    },
    {
      tabName: "identity-providers",
      displayName: "Identity Providers",
      to: () => `/access/identity-providers`,
    },
  ];

  return (
    <div className={classNames("details-subnav", className)}>
      <ul>
        {isAccess ?
          accessConfig.map((link, idx) => {
            const generatedMenuItem = (
              <li
                key={idx}
                className={classNames({
                  "is-active": activeTab === link.tabName
                })}>
                <Link to={link.to()}>
                  {link.displayName}
                </Link>
              </li>
            );
            return generatedMenuItem;
          }).filter(Boolean)
          :
          subNavConfig.map((link, idx) => {
            let hasBadge = false;
            if (link.hasBadge) {
              hasBadge = link.hasBadge(watch || {});
            }
            const generatedMenuItem = (
              <li
                key={idx}
                className={classNames({
                  "is-active": activeTab === link.tabName
                })}>
                <Link to={link.to(slug, kotsSequence)}>
                  {link.displayName} {hasBadge && <span className="subnav-badge" />}
                </Link>
              </li>
            );
            if (link.displayRule) {
              return link.displayRule(watch || {}, isVeleroInstalled) && generatedMenuItem;
            }
            return generatedMenuItem;
          }).filter(Boolean)}
      </ul>
    </div>
  );
}

SubNavBar.defaultProps = {
  watch: {}
};

SubNavBar.propTypes = {
  className: PropTypes.string,
  activeTab: PropTypes.string,
  slug: PropTypes.string,
  watch: PropTypes.object
};
