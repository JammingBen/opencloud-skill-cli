# livePhoto

Apple Live Photo metadata. A Live Photo is a pair of files sharing one content
identifier: a still image (HEIC or JPEG) and a short QuickTime video. Unlike a
Motion Photo, the two parts are separate driveItems. The presence of this facet
indicates the item was captured as part of a Live Photo; the counterpart is the
driveItem whose livePhoto facet carries the same contentId. Whether an item is
the still image or the video half follows from its photo/video facets and MIME
type.

There is no public file format specification; the metadata keys are documented
as AVFoundation metadata identifiers:
https://developer.apple.com/documentation/avfoundation/avmetadataidentifier
On the still image they are stored in the Apple maker note, on the video as
com.apple.quicktime.* QuickTime metadata keys.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `contentId` | string | Yes | Identifier (UUID) shared by the still image and the paired video of one
Live Photo. Read-only.
 |
| `stillImageTimeUs` | integer (int64) | No | Time in microseconds within the paired video at which the still image was
captured. The value comes from the video file only: it is the presentation
time of the com.apple.quicktime.still-image-time timed metadata sample in
the QuickTime movie, so it is only present on the video half. The still
image carries no reliable equivalent (the Apple maker note tag 0x0017,
historically documented as a derivable video index, does not encode a
usable time on current iOS versions). If absent, readers should use a
timestamp near the middle of the video track. Read-only.
 |
| `auto` | boolean | No | True if the device decided automatically whether to capture the Live Photo
video (com.apple.quicktime.live-photo.auto). Only present on the video half.
Read-only.
 |
| `vitalityScore` | number (double) | No | Score between 0 and 1 rating how interesting the motion clip is, used by
readers to decide whether to autoplay it
(com.apple.quicktime.live-photo.vitality-score). Only present on the video
half. Read-only.
 |
| `vitalityScoringVersion` | integer (int64) | No | Version of the algorithm that produced vitalityScore
(com.apple.quicktime.live-photo.vitality-scoring-version). Only present on
the video half. Read-only.
 |

