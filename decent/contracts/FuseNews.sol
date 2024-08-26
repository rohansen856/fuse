// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract FuseNewsContract {
    struct FuseNews {
        string id;
        string title;
        string content;
        string createdAt;
        string publisher;
    }

    bool public checkRunning = true;
    FuseNews[] public news;
    event FuseNewsPublish(string id, string title, string content, string createdAt, string publisher);

    function publish(
        string memory id,
        string memory title,
        string memory content,
        string memory createdAt,
        string memory publisher
    ) external {
        news.push(FuseNews(id, title, content, createdAt, publisher));
        emit FuseNewsPublish(id, title, content, createdAt, publisher);
    }

    function getAllFuseNews() external view returns (FuseNews[] memory) {
        return news;
    }
}
