// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IBase64 {

    /// @dev Encodes the input data into a base64 string
    function encode(bytes memory _data) external pure returns (string memory);

    /// @dev Encodes the input data into a URL-safe base64 string
    function encodeURL(bytes memory _data) external pure returns (string memory);

    /// @dev Decodes the input base64 string into bytes
    function decode(string memory _data) external pure returns (bytes memory);

    /// @dev Decodes the input URL-safe base64 string into bytes
    function decodeURL(string memory _data) external pure returns (bytes memory);

}
